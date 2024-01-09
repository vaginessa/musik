let nextSong = null; // Global variable to store the next song
let isPlaying = false; // Global variable to track if audio is playing

$(document).ready(function() {
    // Fetch and play the initial song
    fetchAndPlayInitialSong();

    // Add click event listener to the "New Song" button
    $('#newSong').click(fetchNewSong);

    // Add event listeners to the audio player
    addAudioPlayerEventListeners();

    // Event listener for play button
    $('#playButton').click(function() {
        isPlaying = true;
        // Logic to play the audio
    });

    // Event listener for pause button
    $('#pauseButton').click(function() {
        isPlaying = false;
        // Logic to pause the audio
    });
});

function fetchAndPlayInitialSong() {
    $.ajax({
        url: '/api/random-song', // Replace with your API endpoint
        type: 'GET',
        success: function(data) {
            updatePlayerAndArtwork(data); // Play the fetched song
            fetchNextSong(); // Preload the next song
        },
        error: function(error) {
            console.error('Error fetching initial song:', error);
        }
    });
}

// addAudioPlayerEventListeners();
function addAudioPlayerEventListeners() {
    const audioPlayer = document.getElementById('audioPlayer');
    // Add an event listener for the 'ended' event
    audioPlayer.addEventListener('ended', function() {
        console.log('Song ended. Fetching new song...');
        fetchNewSong(); // Call the function to fetch a new song
    });

    audioPlayer.addEventListener('error', function(e) {
        const error = e.target.error; // Get the MediaError object
        console.error('Audio playback error:', error);
        console.log('Error code:', error.code);
        console.log('Error message:', error.message);
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


}

function fetchNextSong() {
    $.ajax({
        url: '/api/random-song', // Replace with your API endpoint
        type: 'GET',
        success: function(data) {
            // Assuming 'data' contains the song URL, artwork URL, and stream URL
            nextSong = data; // Store the song data for later use
        },
        error: function(error) {
            console.error('Error fetching next song:', error);
        }
    });
}

function fetchNewSong() {
    // Check if there is a preloaded song
    if (nextSong) {
        // Use the preloaded song details
        updatePlayerAndArtwork(nextSong);

        // Fetch the next song in the background
        fetchNextSong();
    } else {
        // Fallback if no song is preloaded
        console.error('No preloaded song available');
    }
}


function updatePlayerAndArtwork(songData) {

    // Update the artwork
    document.getElementById('coverImg').src = songData.artworkUrl;
    // update background image
    document.querySelector('.cover-bg img').src = songData.artworkUrl;
    document.getElementById('songSource').href = songData.songUrl;


    // Update the audio stream
    var audioPlayer = document.getElementById('audioPlayer');


    if (Hls.isSupported()) {
        var hls = new Hls();
        hls.loadSource(songData.streamUrl);
        hls.attachMedia(audioPlayer);
        hls.on(Hls.Events.MEDIA_ATTACHED, function () {
            audioPlayer.muted = false;
            if (isPlaying) {
                audioPlayer.play(); // Play only if isPlaying is true
            }
        });
    } else if (audioPlayer.canPlayType('application/vnd.apple.mpegurl')) {
        audioPlayer.src = songData.streamUrl;
        audioPlayer.addEventListener('loadedmetadata', function() {
            if (isPlaying) {
                audioPlayer.play(); // Play only if isPlaying is true
            }
        });
    }

}

