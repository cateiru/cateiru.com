import {useRouter} from 'next/router';
import React from 'react';
import {useRecoilState} from 'recoil';
import {UserState} from '../../utils/state/atoms';
import {Verify} from '../../utils/verify';

export interface ReturnValues {
  show: boolean;
}

export const useRequire = (
  isRedirect: boolean,
  redirectPath?: string
): ReturnValues => {
  const [user, setUser] = useRecoilState(UserState);
  const router = useRouter();
  const [show, setShow] = React.useState(false);

  React.useEffect(() => {
    const f = async () => {
      if (user) {
        setShow(true);
        return;
      }

      const verify = new Verify(
        isRedirect,
        u => setUser(u),
        router,
        redirectPath
      );

      await verify.exec(() => {
        setShow(true);
      });
    };
    f();
  }, []);

  React.useEffect(() => {
    if (user === null) {
      setShow(false);
    }
  }, [user]);

  return {
    show,
  };
};
