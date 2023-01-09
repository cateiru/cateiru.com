import {ColorModeScript} from '@chakra-ui/react';
import Document, {Html, Head, Main, NextScript} from 'next/document';
import {GA_TRACKING_ID} from '../utils/ga/gtag';

export default class MyDocument extends Document {
  render(): JSX.Element {
    return (
      <Html lang="ja">
        <Head>
          {GA_TRACKING_ID && (
            <>
              <script
                async
                src={`https://www.googletagmanager.com/gtag/js?id=${GA_TRACKING_ID}`}
              />
              <script
                dangerouslySetInnerHTML={{
                  __html: `
             window.dataLayer = window.dataLayer || [];
             function gtag(){dataLayer.push(arguments);}
             gtag('js', new Date());
             gtag('config', '${GA_TRACKING_ID}', {
               page_path: window.location.pathname,
             });
               `,
                }}
              />
            </>
          )}
          {/* ogp */}
          <meta name="description" content="cateiru's page" />
          <meta property="og:title" content="cateiru.com" />
          <meta property="og:description" content="cateiru's page" />
          <meta property="og:type" content="website" />
          <meta property="og:url" content="https://cateiru.com" />
          <meta property="og:image" content="https://cateiru.com/ogp.png" />
          {/* Twitter ogp */}
          <meta name="twitter:card" content="summary_large_image" />
          <meta name="twitter:title" content="cateiru.com" />
          <meta name="twitter:description" content="cateiru's page" />
          <meta name="twitter:image" content="https://cateiru.com/ogp.png" />
          <meta name="twitter:site" content="@cateiru" />
          <meta name="twitter:creator" content="@cateiru" />

          <meta name="referrer" content="strict-origin-when-cross-origin" />

          {/* favicons */}
          <link
            rel="apple-touch-icon"
            sizes="180x180"
            href="/favicons/apple-touch-icon.png"
          />
          <link
            rel="icon"
            type="image/png"
            sizes="32x32"
            href="/favicons/favicon-32x32.png"
          />
          <link
            rel="icon"
            type="image/png"
            sizes="16x16"
            href="/favicons/favicon-16x16.png"
          />
          <link rel="manifest" href="/favicons/site.webmanifest" />
          <link
            rel="mask-icon"
            href="/favicons/safari-pinned-tab.svg"
            color="#572bcf"
          />
          <meta name="msapplication-TileColor" content="#572bcf" />
          <meta name="theme-color" content="#ffffff" />

          <script
            src="https://cdnjs.cloudflare.com/ajax/libs/three.js/r134/three.min.js"
            async
          />
          <script
            src="https://cdn.jsdelivr.net/npm/vanta/dist/vanta.waves.min.js"
            async
          />
        </Head>
        <body>
          <ColorModeScript initialColorMode="system" />
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}
