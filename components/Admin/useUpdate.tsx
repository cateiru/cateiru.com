import {useToast} from '@chakra-ui/react';
import React from 'react';
import type {
  FieldValues,
  FormState,
  UseFormHandleSubmit,
  UseFormSetValue,
  UseFormRegister,
} from 'react-hook-form';
import {useForm} from 'react-hook-form';
import {api} from '../../utils/api';
import {MultiLang} from '../../utils/config/lang';
import {UpdateHandler} from './useList';

interface Args<T, V extends FieldValues> {
  formFunc(d: V, id: number): [FormData, boolean];
  parse: (r: any) => T;
  update: UpdateHandler<T>;
  onClose: () => void;
  target: T | undefined;
  targetId: (t: T) => number;
  setValues: (target: T, setValue: UseFormSetValue<V>) => void;
  path: string;
  deleteIdName: string;
  convertLang: (e: MultiLang) => string;
}

interface ReturnValues<V extends FieldValues> {
  onSubmit: (d: V) => Promise<() => void>;
  handleSubmit: UseFormHandleSubmit<V>;
  register: UseFormRegister<V>;
  formState: FormState<V>;
  wrapperOnClose: () => void;
  onDelete: () => void;
}

/**
 *
 * @param {Args} a - arg
 * @returns {ReturnValues} - return
 */
export function useUpdate<T, V extends FieldValues>(
  a: Args<T, V>
): ReturnValues<V> {
  const {handleSubmit, register, reset, setValue, formState} = useForm<V>();
  const toast = useToast();

  React.useEffect(() => {
    if (!a.target) {
      return;
    }
    a.setValues(a.target, setValue);
  }, [a.target]);

  const wrapperOnClose = () => {
    reset();
    a.onClose();
  };

  const onSubmit = async (d: V) => {
    if (!a.target) {
      return () => {};
    }
    const id = a.targetId(a.target);

    const [form, changed] = a.formFunc(d, id);

    if (!changed) {
      return () => {};
    }

    const res = await fetch(api(a.path), {
      method: 'PUT',
      credentials: 'include',
      mode: 'cors',
      body: form,
    });

    if (res.ok) {
      const responseData = a.parse(await res.json());
      a.update(responseData, 'upd');
      toast({
        status: 'success',
        title: a.convertLang({ja: '更新しました', en: 'Updated Success'}),
      });
    } else {
      toast({
        status: 'error',
        title: (await res.json()).message,
      });
    }

    wrapperOnClose();

    return () => {};
  };

  const onDelete = () => {
    const f = async () => {
      if (!a.target) {
        return;
      }

      const res = await fetch(
        api(`${a.path}?${a.deleteIdName}=${a.targetId(a.target)}`),
        {
          method: 'DELETE',
          credentials: 'include',
          mode: 'cors',
        }
      );

      if (res.ok) {
        a.update(a.target, 'del');
        wrapperOnClose();

        toast({
          status: 'success',
          title: a.convertLang({ja: '削除しました', en: 'Deleted Success'}),
        });
      } else {
        toast({
          status: 'error',
          title: (await res.json()).message,
        });
      }
    };

    f();
  };

  return {
    onSubmit,
    wrapperOnClose,
    onDelete,
    handleSubmit,
    register,
    formState,
  };
}
