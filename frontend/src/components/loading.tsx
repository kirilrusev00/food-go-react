import React from 'react';
import {
  Box, CircularProgress, BoxProps, makeStyles,
} from '@material-ui/core';

const useStyles = makeStyles(() => ({
  error: {
    letterSpacing: '0.25px',
    overflowWrap: 'break-word',
    wordWrap: 'break-word',
    fontSize: '20px',
    width: '100%',
    marginBottom: '8px',
  },
}));

export interface LoadingProps extends BoxProps {
  /* eslint-disable-next-line @typescript-eslint/no-explicit-any */
  error: any;
  loading: boolean;
  iconSize: number;
}

export function Loading({
  loading, error, iconSize, children, ...rest
}: LoadingProps) {
  const classes = useStyles();

  if (loading) {
    return (
      <Box textAlign="center" {...rest}>
        <CircularProgress style={{ width: iconSize, height: iconSize }} />
      </Box>
    );
  }

  if (error) {
    return (
      <Box textAlign="center" color="red" className={classes.error} {...rest}>
        {error.name === 'UnprocessableImageError' ? 'This image contains no qr code encoding a food!' : error.message}
      </Box>
    );
  }
  /* eslint-disable-next-line react/jsx-no-useless-fragment */
  return <>{children}</>;
}

Loading.defaultProps = {
  iconSize: 40,
};
