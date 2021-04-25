import { useQuery } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';
import Environment from '../types/environment';

interface EnvironmentsPayload {
  data: Environment[];
  count: number;
}

const useEnvironments = () =>
  useQuery<EnvironmentsPayload, AxiosError>(
    'fetchEnvironments',
    async () => (await client.get('/api/environments')).data,
  );

export default useEnvironments;
