import React, { useState, FC, useEffect } from 'react';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import { Link as RouterLink } from 'react-router-dom';
import FormHelperText from '@material-ui/core/FormHelperText';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { DefaultApi } from '../../api/apis';
import {
  EntUser,
} from '../../api/models/';

// function Copyright() {
//   return (
//     <Typography variant="body2" color="textSecondary" align="center">
//       {'Copyright © '}
//       <Link color="inherit" href="https://material-ui.com/">
//         Your Website
//       </Link>{' '}
//       {new Date().getFullYear()}
//       {'.'}
//     </Typography>
//   );
// }

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%',
    marginTop: theme.spacing(0),
  },
  submit: {
    margin: theme.spacing(1, 0, 2),
  },
  formControl: {
    margin: theme.spacing(3),
  },
  button: {
    margin: theme.spacing(1, 1, 0, 0),
  },
}));

const SignIn: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [invalidPassOrUser, setInvalid] = useState(' ');
  const [users, setUsers] = useState<EntUser[]>(Array);
  const [email, setEmail] = useState(String);
  const [password, setPassword] = useState(String);
  const [value, setValue] = React.useState('');
  const [error, setError] = React.useState(true);
  const [path, setPath] = useState(String);
  const [helperText, setHelperText] = React.useState('Choose wisely');
  useEffect(() => {
    const getUsers = async () => {
      const res = await api.listUser();
      setLoading(false);
      setUsers(res);
    };
    getUsers();

    const resetLocalStorage = async () => {
      setLoading(false);
      localStorage.setItem("userID", JSON.stringify(null));
      localStorage.setItem("role", JSON.stringify(null));
      localStorage.setItem("userName", JSON.stringify(null));
    }
    resetLocalStorage();

  }, [loading]);

  const CheckUser = async () => {
    users.map((item: any) => {
      
      if ((item.uSEREMAIL == email) && (item.pASSWORD == password)) {

        localStorage.setItem("userName", JSON.stringify(item.uSERNAME));
        localStorage.setItem("userID", JSON.stringify(item.id));
        localStorage.setItem("role", JSON.stringify(item.edges.position.rOLENAME));
        localStorage.setItem("valid", JSON.stringify(null));
        if ((item.edges.position.rOLENAME === value) && (value === "Librarian")){
          window.location.href = "/"
          setPath("/")
        } else if((item.edges.position.rOLENAME === value) && (value === "Library Member")){
          window.location.href = "/"
          setPath("/")
        }
        else {
          setPath("/SignIn")
          localStorage.setItem("valid", JSON.stringify("Invalid Username or Password"));
          setInvalid(''+localStorage.getItem("valid"))
        }
      }
    })
  };

  const handleRadioChange = (event: any) => {
    setValue(event.target.value);
    setHelperText('  ');
    setError(false);
  };

  const EmailHandleChange = (event: any) => {
    setEmail(event.target.value as string);
  };

  const PasswordHandleChange = (event: any) => {
    setPassword(event.target.value as string);
  };
  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <Typography component="h1" variant="subtitle1" color="error">
        {invalidPassOrUser}
        </Typography>
        <form className={classes.form} noValidate>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            value={email}
            onChange={EmailHandleChange}
            autoComplete="email"
            autoFocus
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            value={password}
            onChange={PasswordHandleChange}
            autoComplete="current-password"

          />
          <RadioGroup aria-label="quiz" name="quiz" value={value} onChange={handleRadioChange}>
                <div>
                  &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
                  <FormControlLabel value="Librarian" control={<Radio />} label="Librarian" />
                  &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
                  <FormControlLabel value="Library Member" control={<Radio />} label="Library Member" />
                </div>
              </RadioGroup>
              <FormHelperText>{helperText}</FormHelperText>
          <Button
            disabled = {error}
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={() => {
              CheckUser();
            }}
            component={RouterLink}
            to={path}
          >
            Sign In
          </Button>

        </form>
      </div>

    </Container>
  );
};

export default SignIn;