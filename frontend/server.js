
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
        // const songUrl = songResponse.data.artworkUrl
        const songData = songResponse.data;
        // Extract each value
        const streamUrl = songData.stream_url;
       const artworkUrl = songData.artwork_url;
        const songTitle = songData.song_title;
        const trackUrl = songData.track_url;


        // Fetch artwork and stream URLs for the song
        const artworkResponse = await axios.get(`http://localhost:8080/api/v1/music/artwork?url=${trackUrl}`);
        // const streamResponse = await axios.get(`http://localhost:8080/api/v1/music/stream?url=${songUrl}`);

        console.log("Track URL:", trackUrl);
        console.log("Artwork URL:", artworkUrl);
        console.log("Stream URL:", streamUrl);

        res.json({
            songUrl: trackUrl,
            artworkUrl: artworkResponse.data,
            streamUrl: streamUrl,
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