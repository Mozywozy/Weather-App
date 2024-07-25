document.getElementById('search-form').addEventListener('submit', function(e) {
    e.preventDefault();
    const city = document.getElementById('city-input').value;
    fetch(`/weather?city=${city}`, {
        headers: {
            'X-Requested-With': 'XMLHttpRequest'
        }
    })
        .then(response => response.json())
        .then(data => {
            document.querySelector('.city').textContent = data.city;
            document.querySelector('.temp span:nth-child(1)').textContent = data.temp;
            document.querySelector('.description').textContent = data.description;
            document.querySelector('.icon img').src = `http://openweathermap.org/img/wn/${data.icon}@2x.png`;
        })
        .catch(error => console.error('Error:', error));
});

function updateTime() {
    const now = new Date();
    const days = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
    const day = days[now.getDay()];
    const hours = now.getHours().toString().padStart(2, '0');
    const minutes = now.getMinutes().toString().padStart(2, '0');
    const seconds = now.getSeconds().toString().padStart(2, '0');
    
    document.getElementById('day').innerText = day;
    document.getElementById('clock').innerText = `${hours}:${minutes}:${seconds}`;
}

function addSubmitAnimation() {
    const form = document.querySelector('.search-box form');
    form.addEventListener('submit', function() {
        document.querySelector('.weather-card').style.animation = 'fadeInUp 1s forwards';
    });
}

document.addEventListener('DOMContentLoaded', function() {
    updateTime();
    setInterval(updateTime, 1000);
    addSubmitAnimation();
});

