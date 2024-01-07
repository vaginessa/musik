$(document).ready(function() {
    // Fetch and update the artwork URL
    fetchArtwork();
    fetchAudioStream();

    // Define function to fetch artwork
    function fetchArtwork() {
        $.ajax({
            url: '/api/artwork',
            type: 'GET',
            success: function(data) {
                console.log("success");
                console.log("Received data:", data); // Log the entire data object
                $('#coverImg').attr('src', data.artworkUrl); // Update the src attribute of the image
                $('.cover-bg img').attr('src', data.artworkUrl); // Update the src attribute of the image

            },
            error: function(error) {
                console.error('Error fetching artwork:', error);
            }
        });
    }

    // Define function to fetch audio stream
    function fetchAudioStream() {
        $.ajax({
            url: '/api/stream', // Endpoint in your Node.js server
            type: 'GET',
            success: function(data) {
                var audioUrl = data.streamUrl;
                console.log("Stream URL: ", data.streamUrl)// Adjust according to the response structure
                var audio = document.getElementById('audioPlayer');

                if (Hls.isSupported()) {
                    var hls = new Hls();
                    hls.loadSource(audioUrl);
                    hls.attachMedia(audio);
                    hls.on(Hls.Events.MEDIA_ATTACHED, function () {
                        audio.muted = false; // Mute the audio to avoid autoplay issues
                        audio.play();
                    });
                } else if (audio.canPlayType('application/vnd.apple.mpegurl')) {
                    audio.src = audioUrl;
                    audio.addEventListener('loadedmetadata', function() {
                        audio.play();
                    });
                }
            },
            error: function(error) {
                console.error('Error fetching audio stream:', error);
            }
        });
    }

    // ... other scripts ...
});