import * as React from 'react';
import { useHistory, useParams } from 'react-router-dom';
import { DataGrid, GridColDef, GridValueGetterParams } from '@material-ui/data-grid';
import Application from '../../types/application';

const columns: GridColDef[] = [
  { field: 'name', headerName: 'Name' },
  {
    field: 'image',
    headerName: 'Image',
    valueGetter: (params: GridValueGetterParams) => {
      const image = params.getValue('docker_image_name');
      const tag = params.getValue('docker_image_tag');
      return `${image}:${!tag || tag === '' ? 'latest': tag}`;
    },
    flex: 0.1,
  },
  {
    field: 'port',
    headerName: 'Port',
    flex: 0.1,
  },
  {
    field: 'container_port',
    headerName: 'Container Port',
    flex: 0.1,
  },
];

const ApplicationsTable = ({ apps }: { apps: Application[] }) => {
  const history = useHistory();
  const { envId } = useParams();
  if (!apps) {
    return null;
  }
  return (
    <DataGrid
        rows={apps}
        columns={columns}
        pageSize={10}
        autoHeight
        disableColumnSelector
        disableSelectionOnClick
        onRowClick={(row) => {
          history.push(`/environments/${envId}/applications/${row.id}`);
        }}
      />
  );
}

export default ApplicationsTable;
