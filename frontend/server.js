
const express = require('express');
const path = require('path');
const app = express();

// Serve static files from a directory named 'public'
app.use(express.static(path.join(__dirname, 'public')));

// Serve your index.html file at the root route
app.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, 'public/index.html'));
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, '0.0.0.0', () => {
    console.log(`Server running on port ${PORT}`);
});