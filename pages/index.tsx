import type {GetStaticProps, NextPage} from 'next';
import React from 'react';
import {Head} from '../components/Common/Head';
import Index from '../components/Index';
import {consolePublic} from '../utils/parse';
import {getPublicProfile} from '../utils/public';
import {Public} from '../utils/types';

const CACHE_TIME = 60; // 1 min

type Props = {
  profile?: Public;
  error?: string;
};

const Home: NextPage<Props> = props => {
  React.useEffect(() => {
    if (props.profile && typeof window !== 'undefined') {
      consolePublic(props.profile);
    }
  }, []);

  return (
    <>
      <Head title={{ja: 'Cateiruのページ', en: "Cateiru's Page"}} />
      {props.profile && <Index profile={props.profile} />}
    </>
  );
};

export const getStaticProps: GetStaticProps<Props> = async () => {
  try {
    const d = await getPublicProfile();
    return {
      props: {
        profile: d,
      },
      revalidate: CACHE_TIME,
    };
  } catch (e) {
    if (e instanceof Error) {
      return {
        props: {
          error: e.message,
        },
        notFound: true,
        revalidate: 1,
      };
    }
  }

  return {
    props: {
      error: 'error',
    },
    revalidate: 1,
    notFound: true,
  };
};

export default Home;
