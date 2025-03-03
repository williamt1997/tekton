document.addEventListener("DOMContentLoaded", () => {
  let sound = document.createElement('audio');
  sound.id = 'audio';
  sound.src = '/assets/backsound.mp3';
  sound.type = 'audio/mp3';
  document.body.appendChild(sound);

  function tryPlayAudio() {
      sound.play().catch(error => {
          console.log('Autoplay blocked. Waiting for user interaction...');
          document.addEventListener('click', tryPlayAudio, { once: true });
      });
  }

  setTimeout(tryPlayAudio, 2700);
});