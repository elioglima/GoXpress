program ToolsDSB;

uses
  Vcl.Forms,
  uPrincipal in 'uPrincipal.pas' {Form1},
  uCNAB in 'uCNAB.pas' {frmCNAB},
  uLoadFile in 'uLoadFile.pas';

{$R *.res}

begin
  Application.Initialize;
  Application.MainFormOnTaskbar := True;
  Application.CreateForm(TForm1, Form1);
  Application.CreateForm(TfrmCNAB, frmCNAB);
  Application.Run;
end.
