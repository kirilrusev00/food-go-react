import {
  AppBar, Box, Link, makeStyles,
  Toolbar, Typography,
} from '@material-ui/core';
import FastfoodIcon from '@material-ui/icons/Fastfood';
import React from 'react';
import { Link as RouterLink } from 'react-router-dom';

const useStyles = makeStyles((theme) => ({
  appBar: {
    top: 0,
    bottom: 'auto',
    backgroundColor: 'white',
    color: 'black',
  },
  offset: theme.mixins.toolbar,
  root: {
    justifyContent: 'space-between',
    paddingLeft: '13.8%',
    paddingRight: '13.8%',
  },
  logo: {
    display: 'flex',
    flexDirection: 'row',
  },
  icon: {
    height: '33px',
    width: '33px',
  },
  title: {
    paddingLeft: theme.spacing(1),
    width: '204px',
  },
}));

export default function Header() {
  const classes = useStyles();

  return (
    <>
      <AppBar position="fixed" className={classes.appBar}>
        <Toolbar className={classes.root}>
          <Box className={classes.logo}>
            <FastfoodIcon className={classes.icon} />
            <Typography variant="h6" className={classes.title}>
              <Link component={RouterLink} to="/" color="primary" underline="none">
                Foods
              </Link>
            </Typography>
          </Box>
        </Toolbar>
      </AppBar>
      <div className={classes.offset} />
    </>
  );
}
