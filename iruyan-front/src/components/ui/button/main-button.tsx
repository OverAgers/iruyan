import Button from '@mui/joy/Button';

type Props = {
  title: string;
  type: 'submit' | 'button' | 'reset' | undefined;
  fullWidth?: boolean;
  maxWidth?: string;
  width?: string;
  onClick?: () => void;
  disabled?: boolean;
  component: 'a' | 'button' | 'div' | 'span' | 'label' | 'input' | 'select' | 'textarea';
};

export default function MainButton(Props: Props) {
  return (
    <Button
      component={Props.component}
      type={Props.type}
      size="lg"
      fullWidth={Props.fullWidth}
      onClick={Props.onClick}
      disabled={Props.disabled}
      sx={{
        backgroundColor: '#7A8764',
        fontWeight: 'bold',
        p: 2,
        maxWidth: Props.maxWidth,
        width: Props.width,
        '&:active': {
          backgroundColor: '#353A2B',
        },
      }}
    >
      {Props.title}
    </Button>
  );
}
