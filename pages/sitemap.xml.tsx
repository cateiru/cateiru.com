import {GetServerSidePropsContext} from 'next';
import {streamToPromise} from 'sitemap';
import {generateSitemap} from '../utils/sitemap';

export const getServerSideProps = async ({res}: GetServerSidePropsContext) => {
  try {
    res.setHeader('Content-Type', 'application/xml');
    res.setHeader('Content-Encoding', 'gzip');
    res.setHeader('Cache-Control', 's-maxage=86400, stale-while-revalidate'); // 24時間のキャッシュ

    const pipeline = await generateSitemap();
    const data = await streamToPromise(pipeline);
    res.end(data);
    pipeline.pipe(res);
    res.statusCode = 200;
  } catch (e) {
    if (e instanceof Error) {
      res.statusCode = 404;
      res.end();
      return;
    }
  }

  return {
    props: {},
  };
};

const Page = () => null;
export default Page;
