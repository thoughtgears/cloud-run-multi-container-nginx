import express from 'express';

const app = express()
const port = process.env.PORT || 8080;

app.use(express.json());

app.get('/ping', (req, res) => {
    res.status(200).json({
        message: 'pong',
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