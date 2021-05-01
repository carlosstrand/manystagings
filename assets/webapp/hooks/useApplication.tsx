import { useQuery } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';
import Application from '../types/application';

const useApplication = (appId: string) =>
  useQuery<Application, AxiosError>(
    ['fetchApplication', appId],
    async () => (await client.get(`/api/applications/${appId}`)).data,
    {
      enabled: !!appId,
    }
  );

export default useApplication;
