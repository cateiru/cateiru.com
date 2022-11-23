import {GetStaticPaths, GetStaticProps, NextPage} from 'next';
import {useRouter} from 'next/router';
import React from 'react';
import {Head} from '../../components/Common/Head';
import {ProductPage} from '../../components/Products/ProductPage';
import {getPublicProduct} from '../../utils/public';
import {PublicProduct} from '../../utils/types';

const CACHE_TIME = 86400; // 1 day

type Props = {
  product?: PublicProduct;
};

const Product: NextPage<Props> = props => {
  const router = useRouter();

  if (router.isFallback && !props.product) return <></>;

  return (
    <>
      <Head
        title={{
          ja: props.product?.name_ja ?? '',
          en: props.product?.name ?? '',
        }}
      />
      {props.product && <ProductPage product={props.product} />}
    </>
  );
};

export const getStaticProps: GetStaticProps<
  Props,
  {id: string}
> = async context => {
  const id = context.params?.id;

  if (typeof id === 'undefined') {
    return {
      notFound: true,
    };
  }

  try {
    const d = await getPublicProduct(parseInt(id));
    return {
      props: {
        product: d,
      },
      revalidate: CACHE_TIME,
    };
  } catch (e) {
    if (e instanceof Error) {
      return {
        notFound: true,
      };
    }
  }

  return {
    notFound: true,
  };
};

export const getStaticPaths: GetStaticPaths = async () => ({
  paths: [],
  fallback: true,
});

export default Product;
