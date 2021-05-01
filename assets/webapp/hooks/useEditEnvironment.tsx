import { useMutation, UseMutationOptions } from 'react-query';
import { AxiosError } from 'axios';
import client from '../client';
import Environment from '../types/environment';

interface EditEnvironmentPayload {
  name: string;
};

const useEditEnvironment = (opts: UseMutationOptions<EditEnvironmentPayload, AxiosError, Environment>) =>
  useMutation<EditEnvironmentPayload, AxiosError, Environment>(
    'editEnvironment',
    async (input: Environment) => client.put<Environment, EditEnvironmentPayload>(`/api/environments/${input.id}`, input),
    opts
  );

export default useEditEnvironment;
