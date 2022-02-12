import React from 'react';
import {
  AppBar, Typography, Toolbar, makeStyles,
} from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  appBar: {
    top: 'auto',
    bottom: 0,
    color: 'black',
    backgroundColor: 'white',
  },
  root: {
    paddingLeft: '13.8%',
    paddingRight: '13.8%',
  },
  offset: theme.mixins.toolbar,
}));

export default function Footer() {
  const classes = useStyles();

  return (
    <>
      <AppBar position="fixed" color="primary" className={classes.appBar}>
        <Toolbar className={classes.root}>
          <Typography>
            © 2022 Food Data
          </Typography>
        </Toolbar>
      </AppBar>
      <div className={classes.offset} />
    </>
  );
}
