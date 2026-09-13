const puppeteer = require('puppeteer');
const fs = require('fs');

(async () => {
    const browser = await puppeteer.launch({
        headless: 'new',
        args: ['--no-sandbox', '--disable-setuid-sandbox']
    });
    const page = await browser.newPage();

    let jsonSpec;
    page.on('response', async (res) => {
        try {
            const url = res.url();
            const ct = (res.headers()['content-type'] || '').toLowerCase();
            if (url.includes('swagger') || url.endsWith('.json') || ct.includes('application/json')) {
                const text = await res.text();
                if (text && text.trim().startsWith('{')) jsonSpec = text;
            }
        } catch (e) { }
    });

    await page.goto('https://developer.statuspage.io/', { waitUntil: 'networkidle2' });

    // Cliquer le lien si présent pour forcer la génération du blob
    const link = await page.$('a[download="swagger.json"]');
    if (!link) throw new Error('Bouton de téléchargement non trouvé.');
    await link.click();

    // Attendre que l'intercepteur capture la spec
    const maxWait = 5000;
    const poll = 100;
    let waited = 0;
    while (!jsonSpec && waited < maxWait) {
        await new Promise((r) => setTimeout(r, poll));
        waited += poll;
    }

    // Fallback simple: lire href et fetch si nécessaire
    if (!jsonSpec) {
        const href = await page.evaluate(() => {
            const l = document.querySelector('a[download="swagger.json"]');
            return l ? l.href : null;
        });
        if (href) {
            if (href.startsWith('data:')) {
                const m = href.match(/^data:([^,]*),(.*)$/);
                if (!m) throw new Error('Data URL malformée.');
                const meta = m[1];
                const data = m[2];
                jsonSpec = meta.endsWith(';base64') ? Buffer.from(data, 'base64').toString('utf8') : decodeURIComponent(data);
            } else {
                jsonSpec = await page.evaluate(async (u) => (await fetch(u)).text(), href);
            }
        }
    }

    if (!jsonSpec) throw new Error('Impossible de récupérer la spec JSON.');

    // Nettoyage et écriture
    if (jsonSpec.charCodeAt(0) === 0xfeff) jsonSpec = jsonSpec.slice(1);
    fs.writeFileSync('developer_statuspage_io.json', jsonSpec.trim());
    console.log('Spécification sauvegardée dans developer_statuspage_io.json');
    await browser.close();
})();