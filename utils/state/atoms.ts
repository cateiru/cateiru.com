import {atom} from 'recoil';
import {User} from '../types';

export const langState = atom({
  key: 'langState',
  default: 'ja',
});

export const UserState = atom<User | null | undefined>({
  key: 'userState',
  default: undefined,
});
