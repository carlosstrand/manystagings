import { useQuery } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';
import Application from '../types/application';

const useApplication = (appId: string) =>
  useQuery<Application, AxiosError>(
    'fetchApplication',
    async () => (await client.get(`/api/applications/${appId}`)).data,
  );

export default useApplication;
