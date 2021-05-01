import { useMutation, UseMutationOptions } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';
import Application from '../types/application';


const useCreateApplication = (opts: UseMutationOptions<Application, AxiosError, Application>) =>
  useMutation<Application, AxiosError, Application>(
    'createApplication',
    async (input: Application) => client.post<Application, Application>('/api/applications', input),
    opts
  );

export default useCreateApplication;
