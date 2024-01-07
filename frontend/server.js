
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

async function getRandomSongDetails(req, res) {
    try {
        // Fetch random song URL
        const songResponse = await axios.get('http://localhost:8080/api/v1/music/random');
        const songUrl = songResponse.data;


        // Fetch artwork and stream URLs for the song
        const artworkResponse = await axios.get(`http://localhost:8080/api/v1/music/artwork?url=${songUrl}`);
        const streamResponse = await axios.get(`http://localhost:8080/api/v1/music/stream?url=${songUrl}`);

        console.log("Song URL:", songUrl);
        console.log("Artwork URL:", artworkResponse.data);
        console.log("Stream URL:", streamResponse.data);

        res.json({
            songUrl: songUrl,
            artworkUrl: artworkResponse.data,
            streamUrl: streamResponse.data
        });
    } catch (error) {
        console.error("Error in getRandomSongDetails:", error);
        res.status(500).send('Error fetching song details');
    }
}

// Route to handle fetching random song details
app.get('/api/random-song', getRandomSongDetails);


const PORT = process.env.PORT || 3000;
app.listen(PORT, '0.0.0.0', () => {
    console.log(`Server running on port ${PORT}`);
});