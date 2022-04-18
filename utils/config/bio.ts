import type {Area} from './area';
import type {MultiLang} from './lang';

export interface Biography {
  primarySchool?: BiographyDetail;
  juniorHighSchool?: BiographyDetail;
  highSchool?: BiographyDetail;
  university?: BiographyDetail;

  works: WorksDetail[];
}

export interface BiographyDetail {
  name: MultiLang;
  area: Area;

  admission: Date;
  graduation?: Date;
}

export interface WorksDetail extends BiographyDetail {
  occupation: MultiLang;
}
