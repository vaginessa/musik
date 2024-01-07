function fetchArtwork(songId) {
    fetch(`/api/artwork?songId=${songId}`)
        .then(response => response.json())
        .then(data => {
            document.getElementById('artworkContainer').innerHTML = '<img src="' + data.artworkUrl + '" alt="Artwork">';
        })
        .catch(error => console.error('Error:', error));
}

// Call this function with the appropriate songId
fetchArtwork('123');