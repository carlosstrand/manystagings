import { useMutation, UseMutationOptions } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';

interface DeleteEnvironmentInput {
  id: string;
}

interface DeleteEnvironmentPayload {
  status: boolean;
}

const useDeleteEnvironment = (opts: UseMutationOptions<DeleteEnvironmentPayload, AxiosError, DeleteEnvironmentInput>) =>
  useMutation<DeleteEnvironmentPayload, AxiosError, DeleteEnvironmentInput>(
    'deleteEnvironment',
    async (input: DeleteEnvironmentInput) => client.delete<DeleteEnvironmentInput, DeleteEnvironmentPayload>(`/api/environments/${input.id}`),
    opts
  );

export default useDeleteEnvironment;
