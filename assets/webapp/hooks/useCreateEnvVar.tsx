import { useMutation, UseMutationOptions } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';

interface EditEnvVarInput {
  application_id: string;
  key: string;
  value: string;
}

interface AuthPayload {
  token: {
    value: string;
    expiration: string;
  };
}

const useCreateEnvVar = (opts: UseMutationOptions<AuthPayload, AxiosError, EditEnvVarInput>) =>
  useMutation<AuthPayload, AxiosError, EditEnvVarInput>(
    'auth',
    async (input: EditEnvVarInput) => client.post<EditEnvVarInput, AuthPayload>('/api/application_env_vars', {
      application_id: input.application_id,
      key: input.key,
      value: input.value,
    }),
    opts
  );

export default useCreateEnvVar;
