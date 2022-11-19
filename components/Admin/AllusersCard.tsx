import {
  Avatar,
  Center,
  Heading,
  Radio,
  RadioGroup,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
  useToast,
} from '@chakra-ui/react';
import React from 'react';
import {useRecoilState} from 'recoil';
import useSWR from 'swr';
import {api} from '../../utils/api';
import {UserState} from '../../utils/state/atoms';
import {fetcher, SWRError} from '../../utils/swr';
import {AllUsers} from '../../utils/types';
import useLanguage from '../useLanguage';
import {CardFrame} from './CardFrame';

export const AllUsersCard = () => {
  const [lang, convertLang] = useLanguage();
  const {data, error} = useSWR<AllUsers, SWRError>('/users', fetcher);
  const toast = useToast();
  const [user, setUser] = useRecoilState(UserState);

  const [selected, setSelected] = React.useState('');

  React.useEffect(() => {
    if (data) {
      setSelected(String(data.find(v => v.selected)?.id) ?? '');
    }
  }, [data]);

  const changeSelected = (id: string) => {
    const f = async () => {
      const form = new FormData();
      form.append('id', id);

      const res = await fetch(api('/user/select'), {
        method: 'POST',
        credentials: 'include',
        mode: 'cors',
        body: form,
      });

      if (res.ok) {
        setSelected(id);

        if (user) {
          const u = {...user};
          if (String(user?.id) === id) {
            u.selected = true;
          } else {
            u.selected = false;
          }
          setUser(u);
        }
        return;
      }

      toast({
        status: 'error',
        title: (await res.json()).message,
      });
    };

    f();
  };

  return (
    <CardFrame>
      <Heading fontSize="1.2rem" textAlign="center" fontWeight="bold">
        {convertLang({ja: '全ユーザ', en: 'All Users'})}
      </Heading>
      {error ? (
        <Center
          fontSize="1.5rem"
          fontWeight="bold"
          color="red.400"
          textAlign="center"
        >
          {error.message}
        </Center>
      ) : (
        <>
          <RadioGroup onChange={changeSelected} value={selected}>
            <TableContainer mt="2rem">
              <Table variant="simple">
                <Thead>
                  <Tr>
                    <Th></Th>
                    <Th>ID</Th>
                    <Th>{convertLang({ja: 'ユーザID', en: 'User ID'})}</Th>
                    <Th>{convertLang({ja: '選択', en: 'Select'})}</Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {data?.map(u => {
                    return (
                      <Tr key={u.id}>
                        <Td fontWeight="bold">
                          <Avatar src={u.avatar_url} size="sm" />
                        </Td>
                        <Td>{u.id}</Td>
                        <Td>{u.user_id}</Td>
                        <Td>
                          <Radio value={`${u.id}`} />
                        </Td>
                      </Tr>
                    );
                  })}
                </Tbody>
              </Table>
            </TableContainer>
          </RadioGroup>
        </>
      )}
    </CardFrame>
  );
};
