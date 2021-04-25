import { useMutation, UseMutationOptions } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';

interface EditEnvVarInput {
  id: string;
  key: string;
  value: string;
}

interface AuthPayload {
  token: {
    value: string;
    expiration: string;
  };
}

const useEditEnvVar = (opts: UseMutationOptions<AuthPayload, AxiosError, EditEnvVarInput>) =>
  useMutation<AuthPayload, AxiosError, EditEnvVarInput>(
    'auth',
    async (input: EditEnvVarInput) => client.put<EditEnvVarInput, AuthPayload>(`/api/application_env_vars/${input.id}`, {
      key: input.key,
      value: input.value,
    }),
    opts
  );

export default useEditEnvVar;
