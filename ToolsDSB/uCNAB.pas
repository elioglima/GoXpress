unit uCNAB;

interface

uses
  Winapi.Windows, Winapi.Messages, System.SysUtils, System.Variants, System.Classes, Vcl.Graphics,
  Vcl.Controls, Vcl.Forms, Vcl.Dialogs, Vcl.StdCtrls;

type
  TCNABTPBanco = (CaixaEconomica);
  TfrmCNAB = class(TForm)
    Log: TMemo;
    Button1: TButton;
    Label1: TLabel;
    OpenDialog: TOpenDialog;
    procedure Button1Click(Sender: TObject);
  private
    { Private declarations }
    Banco:TCNABTPBanco;
  public
    { Public declarations }
    function GerarArquivoTesteCNAB:Boolean;
  end;

var
  frmCNAB: TfrmCNAB;

implementation

{$R *.dfm}

procedure TfrmCNAB.Button1Click(Sender: TObject);
begin
    Close;
end;

function TfrmCNAB.GerarArquivoTesteCNAB: Boolean;
var FileNameRemessa:string;
begin

    if MessageDlg('Deseja gerar um arquivo teste a partir de uma remessa existente?',
        mtConfirmation, [mbYes, mbNo], 0, mbYes) = mrYes then
    begin
        if (not OpenDialog.Execute) then begin
            MessageDlg('Arquivo n�o informado.', mtInformation, [mbOk], 0, mbOk);
            Exit;
        end;

        FileNameRemessa := OpenDialog.FileName;
    end;

end;

end.