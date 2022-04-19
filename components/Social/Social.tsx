import {
  Center,
  Box,
  Heading,
  Text,
  UnorderedList,
  ListItem,
  Link,
} from '@chakra-ui/react';
import React from 'react';
import {IoCaretDown} from 'react-icons/io5';
import config from '../../utils/config/config';
import {MultiLang} from '../../utils/config/lang';
import useLanguage from '../useLanguage';

const Social: React.FC<{
  next: () => void;
  r: React.MutableRefObject<HTMLDivElement>;
}> = ({next, r}) => {
  const [, convertLang] = useLanguage();

  const element = (name: MultiLang, link: string, noLink?: boolean) => {
    return (
      <ListItem my=".5rem">
        {noLink ? (
          <>
            <Text fontWeight="bold" as="span" mr=".5rem">
              {convertLang(name)}:
            </Text>
            <Text as="span">{link}</Text>
          </>
        ) : (
          <Link href={link} isExternal fontWeight="bold">
            {convertLang(name)}
          </Link>
        )}
      </ListItem>
    );
  };

  const Category: React.FC<{
    title: MultiLang;
    children: React.ReactNode;
  }> = props => {
    return (
      <ListItem my="1rem">
        {convertLang(props.title)}
        <UnorderedList>{props.children}</UnorderedList>
      </ListItem>
    );
  };

  return (
    <Center height="100vh" ref={r}>
      <Box width={{base: '90%', md: '500px'}}>
        <Heading textAlign="center">
          {convertLang({en: 'Links', ja: 'リンク'})}
        </Heading>
        <Box whiteSpace="nowrap" overflowX="auto">
          <UnorderedList mt="1.5rem" paddingLeft=".5rem">
            <Category
              title={{ja: 'ソースコードホスティング', en: 'Code Hosting'}}
            >
              {config.socialLinks.github &&
                element(
                  {ja: 'GitHub', en: 'GitHub'},
                  config.socialLinks.github
                )}
              {config.socialLinks.gitlab &&
                element(
                  {ja: 'GitLab', en: 'GitLab'},
                  config.socialLinks.gitlab
                )}
              {config.socialLinks.bitbucket &&
                element(
                  {ja: 'Bitbucket', en: 'Bitbucket'},
                  config.socialLinks.bitbucket
                )}
            </Category>
            <Category title={{ja: 'SNS', en: 'SNS'}}>
              {config.socialLinks.twitter &&
                element(
                  {ja: 'Twitter', en: 'Twitter'},
                  config.socialLinks.twitter
                )}
              {config.socialLinks.facebook &&
                element(
                  {ja: 'Facebook', en: 'Facebook'},
                  config.socialLinks.facebook
                )}
              {config.socialLinks.instagram &&
                element(
                  {ja: 'Instagram', en: 'Instagram'},
                  config.socialLinks.instagram
                )}
              {config.socialLinks.discord &&
                element(
                  {ja: 'Discord', en: 'Discord'},
                  config.socialLinks.discord,
                  true
                )}
            </Category>
            <Category
              title={{
                ja: '動画共有プラットフォーム',
                en: 'Video sharing platform',
              }}
            >
              {config.socialLinks.tiktok &&
                element(
                  {ja: 'TikTok', en: 'TikTok'},
                  config.socialLinks.tiktok
                )}
              {config.socialLinks.youtube &&
                element(
                  {ja: 'YouTube', en: 'YouTube'},
                  config.socialLinks.youtube
                )}
              {config.socialLinks.niconico &&
                element(
                  {ja: 'ニコニコ動画', en: 'niconico'},
                  config.socialLinks.niconico
                )}
            </Category>
            <Category
              title={{
                ja: 'ブログ',
                en: 'blog',
              }}
            >
              {config.socialLinks.hatenaBlog &&
                element(
                  {ja: 'Hatena Blog', en: 'Hatena Blog'},
                  config.socialLinks.hatenaBlog
                )}
              {config.socialLinks.qiita &&
                element({ja: 'Qiita', en: 'Qiita'}, config.socialLinks.qiita)}
              {config.socialLinks.zenn &&
                element({ja: 'Zenn', en: 'Zenn'}, config.socialLinks.zenn)}
              {config.socialLinks.note &&
                element({ja: 'note', en: 'note'}, config.socialLinks.note)}
            </Category>
            <Category
              title={{
                ja: 'その他',
                en: 'other',
              }}
            >
              {config.socialLinks.hatena &&
                element(
                  {ja: 'はてな', en: 'hatena'},
                  config.socialLinks.hatena
                )}
            </Category>
          </UnorderedList>
        </Box>
        <Center mt="1.5rem" cursor="pointer" onClick={next}>
          <IoCaretDown size="25px" />
        </Center>
      </Box>
    </Center>
  );
};

export default Social;
