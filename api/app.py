from quart import Quart, render_template
from quart_cors import cors

app = Quart(__name__)
app = cors(app, allow_origin="*")


@app.route("/")
async def hello():
    return 'hello'


@app.route("/socials")
async def socials():
    import time
    time.sleep(5)
    socials = [
        dict(
            name="Github",
            url="https://github.com/mattcollie",
            svg="""<svg width="12" height="12" viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M5.35161 12.5613C5.35161 12.6258 5.27742 12.6774 5.18387 12.6774C5.07742 12.6871 5.00323 12.6355 5.00323 12.5613C5.00323 12.4968 5.07742 12.4452 5.17097 12.4452C5.26774 12.4355 5.35161 12.4871 5.35161 12.5613ZM4.34839 12.4161C4.32581 12.4806 4.39032 12.5548 4.4871 12.5742C4.57097 12.6065 4.66774 12.5742 4.6871 12.5097C4.70645 12.4452 4.64516 12.371 4.54839 12.3419C4.46452 12.3194 4.37097 12.3516 4.34839 12.4161ZM5.77419 12.3613C5.68064 12.3839 5.61613 12.4452 5.62581 12.5194C5.63548 12.5839 5.71935 12.6258 5.81613 12.6032C5.90968 12.5806 5.97419 12.5194 5.96452 12.4548C5.95484 12.3935 5.86774 12.3516 5.77419 12.3613ZM7.89677 0C3.42258 0 0 3.39677 0 7.87097C0 11.4484 2.25161 14.5097 5.46774 15.5871C5.88064 15.6613 6.02581 15.4065 6.02581 15.1968C6.02581 14.9968 6.01613 13.8935 6.01613 13.2161C6.01613 13.2161 3.75806 13.7 3.28387 12.2548C3.28387 12.2548 2.91613 11.3161 2.3871 11.0742C2.3871 11.0742 1.64839 10.5677 2.43871 10.5774C2.43871 10.5774 3.24194 10.6419 3.68387 11.4097C4.39032 12.6548 5.57419 12.2968 6.03548 12.0839C6.10968 11.5677 6.31935 11.2097 6.55161 10.9968C4.74839 10.7968 2.92903 10.5355 2.92903 7.43226C2.92903 6.54516 3.17419 6.1 3.69032 5.53226C3.60645 5.32258 3.33226 4.45806 3.77419 3.34194C4.44839 3.13226 6 4.2129 6 4.2129C6.64516 4.03226 7.33871 3.93871 8.02581 3.93871C8.7129 3.93871 9.40645 4.03226 10.0516 4.2129C10.0516 4.2129 11.6032 3.12903 12.2774 3.34194C12.7194 4.46129 12.4452 5.32258 12.3613 5.53226C12.8774 6.10323 13.1935 6.54839 13.1935 7.43226C13.1935 10.5452 11.2935 10.7935 9.49032 10.9968C9.7871 11.2516 10.0387 11.7355 10.0387 12.4935C10.0387 13.5806 10.029 14.9258 10.029 15.1903C10.029 15.4 10.1774 15.6548 10.5871 15.5806C13.8129 14.5097 16 11.4484 16 7.87097C16 3.39677 12.371 0 7.89677 0ZM3.13548 11.1258C3.09355 11.1581 3.10323 11.2323 3.15806 11.2935C3.20968 11.3452 3.28387 11.3677 3.32581 11.3258C3.36774 11.2935 3.35806 11.2194 3.30323 11.1581C3.25161 11.1065 3.17742 11.0839 3.13548 11.1258ZM2.7871 10.8645C2.76452 10.9065 2.79677 10.9581 2.86129 10.9903C2.9129 11.0226 2.97742 11.0129 3 10.9677C3.02258 10.9258 2.99032 10.8742 2.92581 10.8419C2.86129 10.8226 2.80968 10.8323 2.7871 10.8645ZM3.83226 12.0129C3.78065 12.0548 3.8 12.1516 3.87419 12.2129C3.94839 12.2871 4.04194 12.2968 4.08387 12.2452C4.12581 12.2032 4.10645 12.1065 4.04194 12.0452C3.97097 11.971 3.87419 11.9613 3.83226 12.0129ZM3.46452 11.5387C3.4129 11.571 3.4129 11.6548 3.46452 11.729C3.51613 11.8032 3.60323 11.8355 3.64516 11.8032C3.69677 11.7613 3.69677 11.6774 3.64516 11.6032C3.6 11.529 3.51613 11.4968 3.46452 11.5387Z" fill="white"/>
            </svg>"""
        ),
        dict(
            name="Linkedin",
            url="https://www.linkedin.com/in/matt-collecutt",
            svg="""<svg width="12" height="12" viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M14.8571 0H1.13929C0.510714 0 0 0.517857 0 1.15357V14.8464C0 15.4821 0.510714 16 1.13929 16H14.8571C15.4857 16 16 15.4821 16 14.8464V1.15357C16 0.517857 15.4857 0 14.8571 0ZM4.83571 13.7143H2.46429V6.07857H4.83929V13.7143H4.83571ZM3.65 5.03571C2.88929 5.03571 2.275 4.41786 2.275 3.66071C2.275 2.90357 2.88929 2.28571 3.65 2.28571C4.40714 2.28571 5.025 2.90357 5.025 3.66071C5.025 4.42143 4.41071 5.03571 3.65 5.03571ZM13.725 13.7143H11.3536V10C11.3536 9.11429 11.3357 7.975 10.1214 7.975C8.88571 7.975 8.69643 8.93929 8.69643 9.93572V13.7143H6.325V6.07857H8.6V7.12143H8.63214C8.95 6.52143 9.725 5.88929 10.8786 5.88929C13.2786 5.88929 13.725 7.47143 13.725 9.52857V13.7143Z" fill="white"/>
            </svg>"""
        ),
        dict(
            name="Instagram",
            url="https://www.instagram.com/matthewcollecutt",
            svg="""<svg width="12" height="12" viewBox="0 0 448 512" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M224.1 141c-63.6 0-114.9 51.3-114.9 114.9s51.3 114.9 114.9 114.9S339 319.5 339 255.9 287.7 141 224.1 141zm0 189.6c-41.1 0-74.7-33.5-74.7-74.7s33.5-74.7 74.7-74.7 74.7 33.5 74.7 74.7-33.6 74.7-74.7 74.7zm146.4-194.3c0 14.9-12 26.8-26.8 26.8-14.9 0-26.8-12-26.8-26.8s12-26.8 26.8-26.8 26.8 12 26.8 26.8zm76.1 27.2c-1.7-35.9-9.9-67.7-36.2-93.9-26.2-26.2-58-34.4-93.9-36.2-37-2.1-147.9-2.1-184.9 0-35.8 1.7-67.6 9.9-93.9 36.1s-34.4 58-36.2 93.9c-2.1 37-2.1 147.9 0 184.9 1.7 35.9 9.9 67.7 36.2 93.9s58 34.4 93.9 36.2c37 2.1 147.9 2.1 184.9 0 35.9-1.7 67.7-9.9 93.9-36.2 26.2-26.2 34.4-58 36.2-93.9 2.1-37 2.1-147.8 0-184.8zM398.8 388c-7.8 19.6-22.9 34.7-42.6 42.6-29.5 11.7-99.5 9-132.1 9s-102.7 2.6-132.1-9c-19.6-7.8-34.7-22.9-42.6-42.6-11.7-29.5-9-99.5-9-132.1s-2.6-102.7 9-132.1c7.8-19.6 22.9-34.7 42.6-42.6 29.5-11.7 99.5-9 132.1-9s102.7-2.6 132.1 9c19.6 7.8 34.7 22.9 42.6 42.6 11.7 29.5 9 99.5 9 132.1s2.7 102.7-9 132.1z" fill="white"/>
            </svg>"""
        ),
    ]

    return await render_template("components/links/social.html", socials=socials)
