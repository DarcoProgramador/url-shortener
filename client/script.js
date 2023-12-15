
const shortenUrl = async () => {
    const longUrl = document.getElementById('long-url').value;
    if (!longUrl) {
        alert('Please enter a URL');
        return;
    }

    if (!await validarURL(longUrl)) {
        alert('Please enter a valid URL');
        return;
    }

    let dataJson = {
        "original_url": String(longUrl)
    };

    const response = await fetch('https://url-shortener-api-production-c3e4.up.railway.app', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(dataJson),
    });

    const data = await response.json();
    const shortUrl = data.short_url;
    const shortUrlContainer = document.getElementById('short-url-container');
    const shortUrlLink = document.getElementById('short-url').querySelector('a');
    shortUrlLink.href = shortUrl;
    shortUrlLink.textContent = shortUrl;
    shortUrlContainer.style.display = 'block';

    const copyBtn = document.getElementById('copy-btn');
    copyBtn.addEventListener('click', () => {
        navigator.clipboard.writeText(shortUrl).then(() => {
            alert('Copied to clipboard!');
        });
    });
};

const validarURL = async (miurl) => {
    try {
        new URL(miurl);
        return true;

    } catch (err) {
        return false;
    }
}

const shortenBtn = document.getElementById('shorten-btn');
shortenBtn.addEventListener('click', shortenUrl);