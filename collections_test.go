// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rapid

import "testing"

func TestCollectionsWithImpossibleMinSize(t *testing.T) {
	s := createRandomBitStream(t)
	gens := []*Generator{
		MapsOfN(Booleans(), Ints(), 10, -1),
		SlicesOfNDistinct(Ints(), 10, -1, func(i int) int { return i % 5 }),
	}

	for _, g := range gens {
		t.Run(g.String(), func(t *testing.T) {
			_, err := recoverValue(g, s)
			if err == nil || !err.isInvalidData() {
				t.Fatalf("got error %v instead of invalid data", err)
			}
		})
	}
}
