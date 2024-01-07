
const express = require('express');
const path = require('path');
const axios = require('axios');
const app = express();

// Serve static files from a directory named 'public'
app.use(express.static(path.join(__dirname, 'public')));

// Serve your index.html file at the root route
app.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, 'public/index.html'));
});

// http://backend:8080/api/v1/music/artwork?url=https://soundcloud.com/uiceheidd/all-girls-are-same-999-prod-nick-mira
// Example route to fetch artwork
app.get('/api/artwork', async (req, res) => {
    try {
        const response = await axios.get('http://localhost:8080/api/v1/music/artwork?url=https://soundcloud.com/uiceheidd/all-girls-are-same-999-prod-nick-mira');
        console.log("This is the url:" , response.data);
        res.json({ artworkUrl: response.data }); // Modify according to the response structure of your backend
    } catch (error) {
        console.error(error);
        res.status(500).send('Error fetching artwork');
    }
});

app.get('/api/stream', async (req, res) => {
    try {
        const response = await axios.get('http://localhost:8080/api/v1/music/stream?url=https://soundcloud.com/uiceheidd/all-girls-are-same-999-prod-nick-mira'); // Replace with your actual backend URL
        console.log("This is the stream data:" , response.data)
        res.json({ streamUrl: response.data }); // Modify according to the response structure of your backend
    } catch (error) {
        console.error(error);
        res.status(500).send('Error fetching stream URL');
    }
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, '0.0.0.0', () => {
    console.log(`Server running on port ${PORT}`);
});