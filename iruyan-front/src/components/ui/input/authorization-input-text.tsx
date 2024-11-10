import { FormControl, FormHelperText, FormLabel, Input } from '@mui/joy';
import { FieldError } from 'react-hook-form';

type Props = {
  label: string;
  placeholder: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  defaultValue?: string;
  endDecorator?: React.ReactNode;
  error: FieldError | undefined;
  type?: 'text' | 'password';
};

export default function AuthInputText(Props: Props) {
  return (
    <FormControl error={Props.error != null} sx={{ mb: 2 }}>
      <FormLabel
        sx={{ mb: 0.5, fontSize: 'lg', fontWeight: 'bold', color: '#3C2800' }}
      >
        {Props.label}
      </FormLabel>
      <Input
        placeholder={Props.placeholder}
        onChange={Props.onChange}
        type={Props.type}
        defaultValue={Props.defaultValue}
        endDecorator={Props.endDecorator}
        fullWidth={true}
        sx={{
          border: 'none',
          borderRadius: '8px',
          boxShadow: 'none',
        }}
      />
      {Props.error != null ? (
        <FormHelperText>{Props.error.message}</FormHelperText>
      ) : null}
    </FormControl>
  );
}
