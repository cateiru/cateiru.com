import {BrandResources} from '../components/BrandResources/BrandResources';
import {Head} from '../components/Common/Head';

const Brand = () => {
  return (
    <>
      <Head title={{ja: 'ブランドリソース', en: 'Brand Resources'}} />
      <BrandResources />
    </>
  );
};

export default Brand;
