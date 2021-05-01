import { useMutation, UseMutationOptions } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';

interface DeleteApplicationInput {
  id: string;
}

interface DeleteApplicationPayload {
  status: boolean;
}

const useDeleteApplication = (opts: UseMutationOptions<DeleteApplicationPayload, AxiosError, DeleteApplicationInput>) =>
  useMutation<DeleteApplicationPayload, AxiosError, DeleteApplicationInput>(
    'deleteApplication',
    async (input: DeleteApplicationInput) => client.delete<DeleteApplicationInput, DeleteApplicationPayload>(`/api/applications/${input.id}`),
    opts
  );

export default useDeleteApplication;
