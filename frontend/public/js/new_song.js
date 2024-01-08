// Define a function to fetch a new song
function fetchNewSong() {
    $.ajax({
        url: '/api/random-song', // Your endpoint that returns a random song URL
        type: 'GET',
        success: function(data) {
            // Assuming 'data' contains the new song URL and artwork URL
            var newSongUrl = data.songUrl; // Adjust this according to your actual response
            var newArtworkUrl = data.artworkUrl; // Adjust this according to your actual response
            var newStreamURL = data.streamUrl;

            console.log("Received new song URL:", newSongUrl);
            console.log("Received new artwork URL:", newArtworkUrl);
            console.log("Received new stream URL:", newStreamURL);

            // Update the artwork
            document.getElementById('coverImg').src = newArtworkUrl;
            // update background image
            document.querySelector('.cover-bg img').src = newArtworkUrl;
            document.getElementById('songSource').href = newSongUrl;

            // Update the audio player
            var audioPlayer = document.getElementById('audioPlayer');
            if (Hls.isSupported()) {
                var hls = new Hls();
                hls.loadSource(newStreamURL);
                hls.attachMedia(audioPlayer);
                hls.on(Hls.Events.MEDIA_ATTACHED, function () {
                    audioPlayer.muted = false;
                    audioPlayer.play();
                });
            } else if (audioPlayer.canPlayType('application/vnd.apple.mpegurl')) {
                audioPlayer.src = newStreamURL;
                audioPlayer.addEventListener('loadedmetadata', function() {
                    audioPlayer.play();
                });
            }
        },
            error: function(error) {
                console.error('Error fetching new song:', error);
            }
    });
}

// Assign the function to the "New Song" button
document.querySelector('.button.next').onclick = fetchNewSong;
