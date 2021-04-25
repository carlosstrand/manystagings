import { useQuery } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';
import Application from '../types/application';

interface ApplicationsPayload {
  data: Application[];
  count: number;
}

const useApplications = (envID: string) =>
  useQuery<ApplicationsPayload, AxiosError>(
    'fetchApplications',
    async () => (await client.get('/api/applications', {
      params: {
        filter: JSON.stringify({
          where: {
            environment_id: {
              eq: envID,
            },
          },
        }),
      },
    })).data,
  );

export default useApplications;
