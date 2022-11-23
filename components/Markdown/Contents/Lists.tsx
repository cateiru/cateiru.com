/**********************************************************
 * md lists
 *
 * @author Yuto Watanabe <yuto.w51942@gmail.com>
 * @version 1.0.0
 *
 * Copyright (C) 2021 hello-slide
 **********************************************************/
import {UnorderedList, ListItem, OrderedList} from '@chakra-ui/react';
import type {
  LiComponent,
  OrderedListComponent,
  UnorderedListComponent,
} from 'react-markdown/lib/ast-to-react';

export const Ul: UnorderedListComponent = ({children}) => {
  return (
    <UnorderedList paddingLeft="1.75rem" marginY="1rem">
      {children}
    </UnorderedList>
  );
};

export const Ol: OrderedListComponent = ({children}) => {
  return (
    <OrderedList paddingLeft="1.75rem" marginY="1rem">
      {children}
    </OrderedList>
  );
};

export const Li: LiComponent = ({children}) => {
  return <ListItem marginY=".5rem">{children}</ListItem>;
};
