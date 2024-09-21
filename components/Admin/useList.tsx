import { useDisclosure } from "@chakra-ui/react";
import React from "react";
import type { KeyedMutator } from "swr";
import useSWR from "swr";
import { fetcher, SWRError } from "../../utils/swr";

type OperationType = "cre" | "upd" | "del";
export type UpdateHandler<T> = (l: T, type: OperationType) => void;
type Equal<T> = (t: T, v: T) => boolean;

interface Modal {
  isOpen: boolean;
  onOpen: () => void;
  onClose: () => void;
  onToggle: () => void;
  isControlled: boolean;
  getButtonProps: (props?: any) => any;
  getDisclosureProps: (props?: any) => any;
}

interface ReturnValues<T> {
  data: T[] | undefined;
  error: SWRError | undefined;
  mutate: KeyedMutator<T[]>;
  update: UpdateHandler<T>;
  updateValue: T | undefined;
  onUpdate: (v: T) => void;
  closeUpdateModal: () => void;
  createModal: Modal;
  updateModal: Modal;
}

/**
 * @param {string} path - api path
 * @param {Equal} e - equal diff handler
 * @returns {ReturnValues} - tools
 */
export function useList<T = any>(path: string, e: Equal<T>): ReturnValues<T> {
  const { data, error, mutate } = useSWR<T[], SWRError>(path, fetcher);
  const [updateValue, setUpdateValue] = React.useState<T | undefined>(
    undefined,
  );

  const createModal = useDisclosure();
  const updateModal = useDisclosure();

  const update = React.useCallback<UpdateHandler<T>>(
    (l, type) => {
      const f = async () => {
        if (!data) {
          return;
        }

        const d = [...data];
        const i = d.findIndex((v) => e(l, v));

        switch (type) {
          case "cre":
          case "upd":
            if (i !== -1) {
              d[i] = l;
            } else {
              d.push(l);
            }
            mutate(d);
            break;
          case "del":
            mutate(d.slice(i, i + 1));
            break;
        }
      };
      f();
    },
    [data],
  );

  const onUpdate = React.useCallback(
    (v: T) => {
      setUpdateValue(data?.find((t) => e(v, t)));
      updateModal.onOpen();
    },
    [updateModal, data],
  );

  const closeUpdateModal = React.useCallback(() => {
    setUpdateValue(undefined);
    updateModal.onClose();
  }, [updateModal, updateValue]);

  return {
    data,
    error,
    mutate,
    update,
    updateValue,
    onUpdate,
    closeUpdateModal,
    createModal,
    updateModal,
  };
}
