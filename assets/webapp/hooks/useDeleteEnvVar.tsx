import { useMutation, UseMutationOptions } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';

interface DeleteEnvVarInput {
  id: string;
}

interface DeleteEnvVarPayload {
  status: boolean;
}

const useDeleteEnvVar = (opts: UseMutationOptions<DeleteEnvVarPayload, AxiosError, DeleteEnvVarInput>) =>
  useMutation<DeleteEnvVarPayload, AxiosError, DeleteEnvVarInput>(
    'deleteEnvVar',
    async (input: DeleteEnvVarInput) => client.delete<DeleteEnvVarInput, DeleteEnvVarPayload>(`/api/application_env_vars/${input.id}`),
    opts
  );

export default useDeleteEnvVar;
