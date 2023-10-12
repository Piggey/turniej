package main

import (
	"reflect"
	"testing"

	"github.com/slaraz/turniej/gra_go/proto"
)

func Test_najlepszaKarta(t *testing.T) {
	type args struct {
		stanGry *proto.StanGry
	}
	tests := []struct {
		name  string
		args  args
		want  proto.Karta
		want1 bool
	}{
		{
			name: "",
			args: args{
				stanGry: &proto.StanGry{
					TwojKolor: proto.KolorZolwia_PURPLE,
					TwojeKarty: []proto.Karta{
						proto.Karta(proto.Karta_value["P1B"]),
						proto.Karta(proto.Karta_value["R2"]),
						proto.Karta(proto.Karta_value["A1B"]),
					},
				},
			},
			want:  proto.Karta_XX,
			want1: false,
		},
		{
			name: "",
			args: args{
				stanGry: &proto.StanGry{
					TwojKolor: proto.KolorZolwia_PURPLE,
					TwojeKarty: []proto.Karta{
						proto.Karta(proto.Karta_value["P1"]),
						proto.Karta(proto.Karta_value["P2"]),
						proto.Karta(proto.Karta_value["A1"]),
					},
				},
			},
			want:  proto.Karta_P2,
			want1: true,
		},
		{
			name: "",
			args: args{
				stanGry: &proto.StanGry{
					TwojKolor: proto.KolorZolwia_PURPLE,
					TwojeKarty: []proto.Karta{
						proto.Karta(proto.Karta_value["G1"]),
						proto.Karta(proto.Karta_value["R2"]),
						proto.Karta(proto.Karta_value["A1"]),
					},
				},
			},
			want:  proto.Karta_A1,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := najlepszaKarta(tt.args.stanGry)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("najlepszaKarta() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("najlepszaKarta() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
