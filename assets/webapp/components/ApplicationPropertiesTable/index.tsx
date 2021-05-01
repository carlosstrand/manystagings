import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

const useStyles = makeStyles({
  table: {
    minWidth: 650,
  },
});

interface ApplicationPropertiesTableProps {
  rows: {
    key: string;
    value: string;
  }[],
}

export default function ApplicationPropertiesTable(props: ApplicationPropertiesTableProps) {
  const classes = useStyles();
  const { rows } = props;
  return (
    <TableContainer>
      <Table className={classes.table} size="small" aria-label="a dense table">
        <TableBody>
          {rows.map((row) => (
            <TableRow key={row.key}>
              <TableCell align="left" width="300"><strong>{row.key}</strong></TableCell>
              <TableCell align="left">{row.value}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
