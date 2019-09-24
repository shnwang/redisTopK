package topk

import "testing"

type TestInfo struct {
    s string;
}

func (ti TestInfo) Val() int {
    return len(ti.s);
}

func (ti TestInfo) Str() string {
    return ti.s;
}

func TestPrint(t *testing.T) {
    tk := TopKNew(2);
    tk.Insert(TestInfo{"milan"});
    tk.Insert(TestInfo{"kaka"});
    tk.Insert(TestInfo{"ronaldo"});

    tk.Print();
}

func TestInsert(t *testing.T) {
    tk := TopKNew(2);
    tk.Insert(TestInfo{"milan"});
    tk.Insert(TestInfo{"kaka"});
    tk.Insert(TestInfo{"ronaldo"});

    want := 2;
    if tk.Len() != want {
        t.Errorf("topk len error got:%d want:%d", tk.Len(), want);
    }
}

func TestMin(t *testing.T) {
    tk := TopKNew(2);
    tk.Insert(TestInfo{"milan"});
    tk.Insert(TestInfo{"kaka"});
    tk.Insert(TestInfo{"ronaldo"});

    want := "milan";
    d := tk.Min();
    if d.Str() != want {
        t.Errorf("topk min error got:%s want:%s", d.Str(), want);
    }
}

func TestMax(t *testing.T) {
    tk := TopKNew(2);
    tk.Insert(TestInfo{"milan"});
    tk.Insert(TestInfo{"kaka"});
    tk.Insert(TestInfo{"ronaldo"});

    want := "ronaldo";
    d := tk.Max();
    if d.Str() != want {
        t.Errorf("topk min error got:%s want:%s", d.Str(), want);
    }
}


