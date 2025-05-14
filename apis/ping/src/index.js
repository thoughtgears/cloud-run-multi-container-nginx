import express from 'express';
import admin from 'firebase-admin';
import firebaseAuthMiddleware from './middleware/auth.js';

const app = express()
const port = process.env.PORT || 8080;
const projectId = process.env.FIREBASE_PROJECT_ID;

if (!projectId) {
    console.error('FIREBASE_PROJECT_ID environment variable is not set.');
    process.exit(1);
}

admin.initializeApp({
    projectId: projectId,
});

app.use(express.json());

app.get('/ping', firebaseAuthMiddleware, (req, res) => {
    res.status(200).json({
        message: 'pong',
        uid: req.user.uid
    });
})

app.get('/health', (req, res) => {
    res.status(200).json({
        message: 'healthy',
    });
})

app.listen(port, () => {
    console.log(`Ping service is running on port ${port}`);
})