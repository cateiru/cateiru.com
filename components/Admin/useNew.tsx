import {useToast} from '@chakra-ui/react';
import {useForm} from 'react-hook-form';
import type {
  FieldValues,
  FormState,
  UseFormHandleSubmit,
  UseFormRegister,
  UseFormReturn,
} from 'react-hook-form';
import {api} from '../../utils/api';
import {MultiLang} from '../../utils/config/lang';
import {UpdateHandler} from './useList';

interface Args<T, V extends FieldValues> {
  path: string;
  formFunc: (d: V) => FormData;
  convertLang: (e: MultiLang) => string;
  parse: (r: any) => T;
  update: UpdateHandler<T>;
  onClose: () => void;
}

interface ReturnValues<V extends FieldValues> {
  onSubmit: (d: V) => Promise<() => void>;
  handleSubmit: UseFormHandleSubmit<V>;
  register: UseFormRegister<V>;
  formState: FormState<V>;
  wrapperOnClose: () => void;
  methods: UseFormReturn<V, any>;
}

/**
 *
 * @param {Args} a - arg
 * @returns {ReturnValues} - return
 */
export function useNew<T, V extends FieldValues>(
  a: Args<T, V>
): ReturnValues<V> {
  const methods = useForm<V>();
  const {handleSubmit, register, reset, formState} = methods;
  const toast = useToast();

  const wrapperOnClose = () => {
    reset();
    a.onClose();
  };

  const onSubmit = async (d: V) => {
    const form = a.formFunc(d);

    const res = await fetch(api(a.path), {
      method: 'POST',
      credentials: 'include',
      mode: 'cors',
      body: form,
    });

    if (res.ok) {
      toast({
        status: 'success',
        title: a.convertLang({ja: '作成しました', en: 'Success created'}),
      });
      const responseData = a.parse(await res.json());
      a.update(responseData, 'cre');
    } else {
      toast({
        status: 'error',
        title: (await res.json()).message,
      });
    }

    wrapperOnClose();

    return () => {};
  };

  return {onSubmit, handleSubmit, register, formState, wrapperOnClose, methods};
}
