// Copyright 2020 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build ignore
// +build ignore

package shaders

var Time float
var Cursor vec2
var ScreenSize vec2

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	texC := position.xy / ScreenSize.xy
	// mid := vec2(ScreenSize.x/2, ScreenSize.y/2)
	lpos := Cursor.xy / ScreenSize.x
	texC2 := position.xy / ScreenSize.x
	//return imageSrc0At(uv)
	texC = mix(texC2, (texC*2.0-lpos*2.0)*0.2*0.5+lpos, (1.0 / (distance((texC2*2.0-lpos*2.0)*10.0*0.5+lpos, lpos) - 1.0))) //Black hole shader
	resColor := imageSrc0At(texC).rgb
	//return imageSrc0At(texC).rgba
	resColor *= clamp(pow(distance(texC2, lpos), 8.8)*300000000.0, 0.0, 1.0)

	return vec4(resColor, 1.0)
}
