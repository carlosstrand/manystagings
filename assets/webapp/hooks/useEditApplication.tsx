import { useMutation, UseMutationOptions } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';
import Application from '../types/application';

interface EditApplicationPayload {
  token: {
    value: string;
    expiration: string;
  };
}

const useEditApplication = (opts: UseMutationOptions<EditApplicationPayload, AxiosError, Application>) =>
  useMutation<EditApplicationPayload, AxiosError, Application>(
    'editApplication',
    async (input: Application) => client.put<Application, EditApplicationPayload>(`/api/applications/${input.id}`, input),
    opts
  );

export default useEditApplication;
