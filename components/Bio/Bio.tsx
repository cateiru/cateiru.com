import {
  Center,
  Box,
  Heading,
  Flex,
  Switch,
  useColorMode,
  Text,
  UnorderedList,
  ListItem,
  ListIcon,
  Badge,
} from '@chakra-ui/react';
import {Gantt, ViewMode} from 'gantt-task-react';
import React from 'react';
import {IoCaretDown} from 'react-icons/io5';
import {getTasks} from '../../utils/parse';
import {Public} from '../../utils/types';
import useLanguage from '../useLanguage';

const Bio: React.FC<{
  next: () => void;
  r: React.MutableRefObject<HTMLDivElement>;
  data: Public;
}> = ({next, r, data}) => {
  const {lang, convertLang} = useLanguage();
  const [viewMode, setViewMode] = React.useState<ViewMode>(ViewMode.Year);
  const {colorMode} = useColorMode();

  const tasks = React.useMemo(() => {
    return getTasks(data.biographies, lang);
  }, [data, lang]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setViewMode(e.target.checked ? ViewMode.Month : ViewMode.Year);
  };

  return (
    <Center height="100vh" ref={r}>
      <Box width={{base: '90%', md: '700px'}}>
        <Heading textAlign="center">
          {convertLang({en: 'Brief personal record', ja: '略歴'})}
        </Heading>
        {data.biographies.length !== 0 ? (
          <>
            <Box
              boxShadow={
                colorMode === 'light'
                  ? '0px 1px 26px -3px #a0acc0'
                  : '0px 1px 26px -3px #000'
              }
              borderRadius="25"
              mt=".5rem"
            >
              <Flex justifyContent="right" pt=".5rem" pr=".5rem">
                <Switch onChange={handleChange}>
                  {convertLang({ja: '月ごとで表示する', en: 'View by month'})}
                </Switch>
              </Flex>
              <Flex mt=".5rem" whiteSpace="nowrap" overflowX="auto" id="gantt">
                <Gantt
                  tasks={tasks}
                  listCellWidth=""
                  viewMode={viewMode}
                  columnWidth={100}
                  barCornerRadius={5}
                  locale={lang === 'ja' ? 'ja' : 'en'}
                />
              </Flex>
            </Box>
            <Box whiteSpace="nowrap" overflowX="auto">
              <UnorderedList mt="1.5rem" paddingLeft=".5rem">
                {data.biographies.map((v, i) => {
                  const start = new Date(v.join);

                  let end: Date | undefined = undefined;
                  if (v.leave === '0001-01-01T00:00:00Z') {
                    end = new Date();
                  }

                  return (
                    <ListItem my=".5rem" key={i}>
                      <Text as="span" display="inline-block" minW="10ch">
                        {start.getFullYear()}&nbsp;-&nbsp;
                        {end ? end.getFullYear() : ''}
                      </Text>
                      <Text as="span" ml="1rem" fontWeight="bold">
                        {convertLang({ja: v.name_ja, en: v.name})}
                      </Text>
                      <Text as="span" ml=".5rem">
                        ({convertLang({ja: v.address_ja, en: v.address})})
                      </Text>
                      <Text as="span" ml=".5rem">
                        - {convertLang({ja: v.position_ja, en: v.position})}
                      </Text>
                      <Text as="span" ml=".5rem">
                        <Badge
                          colorScheme={v.type === 'univ' ? 'blue' : 'green'}
                        >
                          {v.type === 'univ'
                            ? convertLang({ja: '大学', en: 'UNIV'})
                            : convertLang({ja: '会社', en: 'CORP'})}
                        </Badge>
                      </Text>
                    </ListItem>
                  );
                })}
              </UnorderedList>
            </Box>
          </>
        ) : (
          <Text textAlign="center" mt="1rem" color="yellow.500">
            {convertLang({
              ja: '表示できる略歴がありません。',
              en: 'There is no biography to display.',
            })}
          </Text>
        )}

        <Center mt="1.5rem" cursor="pointer" onClick={next}>
          <IoCaretDown size="25px" />
        </Center>
      </Box>
    </Center>
  );
};

export default Bio;
