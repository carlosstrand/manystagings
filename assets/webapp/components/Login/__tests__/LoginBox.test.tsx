import { QueryClient, QueryClientProvider } from 'react-query';
import { MemoryRouter } from 'react-router-dom';
import renderer from 'react-test-renderer';
import LoginBox from '..';

test('LoginBox: Match Snapshot', () => {
  const queryClient = new QueryClient();
  const component = renderer.create(
    <QueryClientProvider client={queryClient}>
      <MemoryRouter>
        <LoginBox />
      </MemoryRouter>
    </QueryClientProvider>
  );
  expect(component.toJSON()).toMatchSnapshot();
});
