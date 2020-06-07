const functions = require('firebase-functions');
const admin = require('firebase-admin');

admin.initializeApp(functions.config().firebase);

const secrets = admin.database().ref('app/secrets');

// Google OAuth Setup
const { google } = require('googleapis');
const oauth2Client = new google.auth.OAuth2(
    "650593328036-lqnverllqqppdv8qfdphdg8a0umgd2v8.apps.googleusercontent.com", //ClientID
    "C4QZ72aqvNJ9MEd2ieHcsI-T", //ClientSecret
    "http://localhost" //RedirectURL
);

// Enable CORS (Allow ‘Access-Control-Allow-Origin’)
const cors = require('cors')({ origin: true });
const express = require('express');
const cookieParser = require('cookie-parser')();
const app = express();
var uid;

const validateFirebaseIdToken = (req, res, next) => {
    console.log('Check if request is authorized with Firebase ID token');

    if ((!req.headers.authorization || !req.headers.authorization.startsWith('Bearer ')) &&
        !(req.cookies && req.cookies.__session)) {
        console.error('No Firebase ID token was passed as a Bearer token in the Authorization header.',
            'Make sure you authorize your request by providing the following HTTP header:',
            'Authorization: Bearer <Firebase ID Token>',
            'or by passing a "__session" cookie.');
        res.status(403).send('Unauthorized');
        return;
    }

    let idToken;
    if (req.headers.authorization && req.headers.authorization.startsWith('Bearer ')) {
        console.log('Found "Authorization" header');
        // Read the ID Token from the Authorization header.
        idToken = req.headers.authorization.split('Bearer ')[1];
    } else if (req.cookies) {
        console.log('Found "__session" cookie');
        // Read the ID Token from cookie.
        idToken = req.cookies.__session;
    } else {
        // No cookie
        res.status(403).send('Unauthorized');
        return;
    }
    admin.auth().verifyIdToken(idToken).then((decodedIdToken) => {
        console.log('ID Token correctly decoded', decodedIdToken);
        req.user = decodedIdToken;
        uid = decodedIdToken.uid;
        return next();
    }).catch((error) => {
        console.error('Error while verifying Firebase ID token:', error);
        res.status(403).send('Unauthorized');
    });
};

app.use(cors);
app.use(cookieParser);
app.use(validateFirebaseIdToken);

// store client token
app.post('/putToken', (req, res) => {
    reqJson = req.body
    childRef = uid + "/" + reqJson.provider + "/" + reqJson.email.replace(".", "^-^-")
    delete reqJson["provider"]
    delete reqJson["email"]
    secrets.child(childRef).set(reqJson)
    const result = { "status": "success" };
    res.setHeader('Content-Type', 'application/json');
    res.send(result);
});

// generate Access token
app.post('/getAccessToken', (req, res) => {
    reqJson = req.body
    childRef = uid + "/" + reqJson.provider + "/" + reqJson.email.replace(".", "^-^-")
    secrets.child(childRef).once('value').then((snapshot) => {
        // Get Refresh Token
        oauth2Client.setCredentials({
            refresh_token: snapshot.val().refresh_token
        });
        //  Get Access Token
        const repResult = oauth2Client.refreshAccessToken((_, tokens) => {
            res.setHeader('Content-Type', 'application/json');
            return res.send({ "access_token": tokens.access_token });
        });
        return repResult
    }).catch(err => {
        res.setHeader('Content-Type', 'application/json');
        res.send({"error": err.message});
    });
});

// get Google OAuth Client Credentials
app.get('/getClientConfig', (req, res) => {
    const result = {"installed":{"client_id":"650593328036-lqnverllqqppdv8qfdphdg8a0umgd2v8.apps.googleusercontent.com","project_id":"box-app-80870","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"C4QZ72aqvNJ9MEd2ieHcsI-T","redirect_uris":["urn:ietf:wg:oauth:2.0:oob","http://localhost"]}};
    res.setHeader('Content-Type', 'application/json');
    res.send(result);
});

exports.app = functions.https.onRequest(app);