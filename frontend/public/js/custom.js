$(document).ready(function() {

    function fetchSongDetails() {
        $.ajax({
            url: '/api/random-song', // Updated endpoint
            type: 'GET',
            success: function(data) {
                console.log("Received data:", data); // Log the entire data object

                // Update the artwork
                $('#coverImg').attr('src', data.artworkUrl);
                $('.cover-bg img').attr('src', data.artworkUrl);
                // update the song source
                $('#songSource').attr('href', data.songUrl);

                // Update the audio stream
                var audioUrl = data.streamUrl;
                var audio = document.getElementById('audioPlayer');

                if (Hls.isSupported()) {
                    var hls = new Hls();
                    hls.loadSource(audioUrl);
                    hls.attachMedia(audio);
                    hls.on(Hls.Events.MEDIA_ATTACHED, function () {
                        audio.muted = false;
                       // audio.play();
                    });
                } else if (audio.canPlayType('application/vnd.apple.mpegurl')) {
                    audio.src = audioUrl;
                    audio.addEventListener('loadedmetadata', function() {
                       // audio.play();
                    });
                }
            },
            error: function(error) {
                console.error('Error fetching song details:', error);
            }
        });
    }


    // Get the audio player element
    const audioPlayer = document.getElementById('audioPlayer');

    // Add an event listener for the 'ended' event
    audioPlayer.addEventListener('ended', function() {
        console.log('Song ended. Fetching new song...');
        fetchNewSong(); // Call the function to fetch a new song
    });

    audioPlayer.addEventListener('error', function(e) {
        console.error('Audio playback error:', e);
    });

    audioPlayer.addEventListener('stalled', function() {
        console.log('Audio playback stalled');
    });

    audioPlayer.addEventListener('waiting', function() {
        console.log('Waiting for audio data...');
    });

    audioPlayer.addEventListener('suspend', function() {
        console.log('Audio playback suspended');
    });
    // Call the function to fetch song details
    fetchSongDetails();

});








